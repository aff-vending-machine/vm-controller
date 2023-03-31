package idle

import "github.com/rs/zerolog/log"

func (s *Stage) console() {
	log.Info().Str("stage", "idle").Msg("idle stage")
}
