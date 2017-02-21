package linkdata

type (
	// http://schema.org/MusicGroup
	MusicGroup struct {
		Organization
	}

	// http://schema.org/MusicAlbum
	MusicAlbum struct {
		CreativeWork
		Release *MusicRelease `json:"albumRelease"`
		Artist  *MusicGroup   `json:"byArtist"`
	}

	// http://schema.org/MusicPlaylist
	MusicPlaylist struct {
		CreativeWork
		TrackCount uint              `json:"numTracks"`
		Track      []*MusicRecording `json:"track"`
	}

	// http://schema.org/MusicRecording
	MusicRecording struct {
		CreativeWork
	}

	// http://schema.org/MusicComposition
	MusicComposition struct {
		CreativeWork
		Composer   *Person         `json:"composer"`
		Lyrics     *CreativeWork   `json:"lyrics"`
		Lyricist   *Person         `json:"lyricist"`
		MusicalKey string          `json:"musicalKey"`
		RecordedAs *MusicRecording `json:"recordedAs"`
	}

	// http://schema.org/MusicRelease
	MusicRelease struct {
		MusicPlaylist
		CatalogNumber string        `json:"catalogNumber"`
		CreditedTo    *Person       `json:"creditedTo"`
		RecordLabel   *Organization `json:"recordLabel"`
		ReleaseOf     *MusicAlbum   `json:"releaseOf"`
	}
)
