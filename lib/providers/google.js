"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const config_1 = require("../config");
const constants_1 = require("../constants");
const is_1 = require("../is");
const logger_1 = require("../logger");
const googleAPIs = require("googleapis");
const googleAuth = require("google-auth-library");
const scope = {
    drive: {
        READ_WRITE: 'https://www.googleapis.com/auth/drive',
        READ_ONLY: 'https://www.googleapis.com/auth/drive.readonly'
    },
    photo: {
        READ_ONLY: 'https://www.googleapis.com/auth/drive.photos.readonly'
    }
};
const auth = new googleAuth();
const authConfig = config_1.default.google.auth;
const authClient = new auth.OAuth2(authConfig.clientID, authConfig.secret, authConfig.callback);
const authorizationURL = () => authClient.generateAuthUrl({
    access_type: 'offline',
    approval_prompt: 'force',
    scope: scope.drive.READ_ONLY
});
const accessTokenExpired = () => is_1.default.value(authConfig.token.refresh) &&
    (authConfig.token.accessExpiration === null || authConfig.token.accessExpiration < new Date());
const minuteEarlier = (ms) => {
    const d = new Date(ms);
    d.setMinutes(d.getMinutes() - 1);
    return d;
};
const verifyToken = () => new Promise((resolve, reject) => {
    authClient.setCredentials({
        access_token: authConfig.token.access,
        refresh_token: authConfig.token.refresh
    });
    if (accessTokenExpired()) {
        authClient.refreshAccessToken((err, tokens) => {
            if (is_1.default.value(err)) {
                logger_1.default.error('Unable to refresh Google access token: %s', err.message);
                reject(err);
            }
            else {
                logger_1.default.infoIcon('lock_outline', 'Refreshed Google access token');
                authClient.setCredentials(tokens);
                authConfig.token.type = tokens.token_type;
                authConfig.token.access = tokens.access_token;
                authConfig.token.accessExpiration = minuteEarlier(tokens.expiry_date);
                resolve();
            }
        });
    }
    else {
        resolve();
    }
});
const getAccessToken = (code) => new Promise((resolve, reject) => {
    authClient.getToken(code, (err, token) => {
        if (is_1.default.value(err)) {
            reject(err);
        }
        else {
            authClient.credentials = token;
            resolve({
                access: token.access_token,
                refresh: token.refresh_token,
                accessExpiration: minuteEarlier(token.expiry_date)
            });
        }
    });
});
const driveConfig = config_1.default.google.drive;
let _drive = null;
function drive() {
    if (_drive === null) {
        _drive = googleAPIs.drive('v3');
    }
    return _drive;
}
const loadGPX = (post, stream) => verifyToken().then(() => new Promise((resolve, reject) => {
    const options = {
        auth: authClient,
        q: `name = '${post.title}.gpx' and '${driveConfig.tracksFolder}' in parents`
    };
    drive().files.list(options, (err, list) => {
        post.triedTrack = true;
        if (err !== null) {
            logger_1.default.error('Error finding GPX for “%s”: %s', post.title, err.message);
            reject(err);
        }
        else if (!is_1.default.array(list.files) || list.files.length == 0) {
            post.hasTrack = false;
            logger_1.default.warn(`No GPX file found for “${post.title}”`);
            reject();
        }
        else {
            const file = list.files[0];
            let purpose = 'Retrieving';
            let icon = 'save';
            if (is_1.default.value(stream)) {
                purpose = 'Downloading';
                icon = 'file_download';
            }
            logger_1.default.infoIcon(icon, '%s GPX for “%s” (%s)', purpose, post.title, file.id);
            resolve(downloadFile(file.id, post, stream));
        }
    });
}));
const downloadFile = (fileId, post, stream) => verifyToken().then(() => new Promise((resolve, reject) => {
    const options = { fileId, auth: authClient, alt: 'media', timeout: 10000 };
    if (is_1.default.value(stream)) {
        stream.on('finish', resolve);
        drive().files
            .get(options)
            .on('error', reject)
            .on('end', () => { post.hasTrack = true; })
            .on('response', res => {
            res.headers[constants_1.header.content.DISPOSITION.toLowerCase()] = `attachment; filename=${post.key}.gpx`;
            res.headers[constants_1.header.content.TYPE.toLowerCase()] = constants_1.mimeType.GPX;
        })
            .pipe(stream);
    }
    else {
        drive().files
            .get(options, (err, body, response) => {
            if (is_1.default.value(err)) {
                reject(err);
            }
            else {
                post.hasTrack = true;
                resolve(body);
            }
        })
            .on('error', reject);
    }
}));
exports.default = {
    auth: {
        url: authorizationURL,
        client: authClient,
        verify: verifyToken,
        expired: accessTokenExpired,
        getAccessToken,
        isEmpty() { return is_1.default.empty(authConfig.token.access); }
    },
    drive: {
        loadGPX
    }
};
//# sourceMappingURL=google.js.map