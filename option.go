package juejinsign

type Option func(s *Sign)

func Aid(aid int64) Option {
	return func(mod *Sign) {
		mod.aid = aid
	}
}

func UUID(uuid int64) Option {
	return func(mod *Sign) {
		mod.uuid = uuid
	}
}

func Cookie(cookie string) Option {
	return func(mod *Sign) {
		mod.cookie = cookie
	}
}

func MsToken(msToken string) Option {
	return func(mod *Sign) {
		mod.msToken = msToken
	}
}

func Bogus(bogus string) Option {
	return func(mod *Sign) {
		mod.bogus = bogus
	}
}

func Token(token string) Option {
	return func(mod *Sign) {
		mod.token = token
	}
}
