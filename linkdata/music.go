package linkdata

type (
	// http://schema.org/MusicGroup
	MusicGroup struct {
		Organization
	}

	// http://schema.org/MusicAlbum
	MusicAlbum struct {
		CreativeWork
		Release *MusicRelease `json:"albumRelease,omitempty"`
		Artist  *MusicGroup   `json:"byArtist,omitempty"`
	}

	// http://schema.org/MusicPlaylist
	MusicPlaylist struct {
		CreativeWork
		TrackCount uint              `json:"numTracks,omitempty"`
		Track      []*MusicRecording `json:"track,omitempty"`
	}

	// http://schema.org/MusicRecording
	MusicRecording struct {
		CreativeWork
	}

	// http://schema.org/MusicComposition
	MusicComposition struct {
		CreativeWork
		Composer   *Person         `json:"composer,omitempty"`
		Lyrics     *CreativeWork   `json:"lyrics,omitempty"`
		Lyricist   *Person         `json:"lyricist,omitempty"`
		MusicalKey string          `json:"musicalKey,omitempty"`
		RecordedAs *MusicRecording `json:"recordedAs,omitempty"`
	}

	// http://schema.org/MusicRelease
	MusicRelease struct {
		MusicPlaylist
		CatalogNumber string        `json:"catalogNumber,omitempty"`
		CreditedTo    *Person       `json:"creditedTo,omitempty"`
		RecordLabel   *Organization `json:"recordLabel,omitempty"`
		ReleaseOf     *MusicAlbum   `json:"releaseOf,omitempty"`
	}
)
