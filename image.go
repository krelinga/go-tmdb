package tmdb

type BackdropImage string

func (b BackdropImage) GetUrl(configuration *Configuration, size BackdropSize) (string, bool) {
	return getImageUrl(configuration.Images.BackdropSizes, configuration.Images.BaseUrl, size, string(b))
}

func (b BackdropImage) GetSecureUrl(configuration *Configuration, size BackdropSize) (string, bool) {
	return getImageUrl(configuration.Images.BackdropSizes, configuration.Images.SecureBaseUrl, size, string(b))
}

type LogoImage string

func (l LogoImage) GetUrl(configuration *Configuration, size LogoSize) (string, bool) {
	return getImageUrl(configuration.Images.LogoSizes, configuration.Images.BaseUrl, size, string(l))
}

func (l LogoImage) GetSecureUrl(configuration *Configuration, size LogoSize) (string, bool) {
	return getImageUrl(configuration.Images.LogoSizes, configuration.Images.SecureBaseUrl, size, string(l))
}

type PosterImage string

func (p PosterImage) GetUrl(configuration *Configuration, size PosterSize) (string, bool) {
	return getImageUrl(configuration.Images.PosterSizes, configuration.Images.BaseUrl, size, string(p))
}

func (p PosterImage) GetSecureUrl(configuration *Configuration, size PosterSize) (string, bool) {
	return getImageUrl(configuration.Images.PosterSizes, configuration.Images.SecureBaseUrl, size, string(p))
}

type ProfileImage string

func (p ProfileImage) GetUrl(configuration *Configuration, size ProfileSize) (string, bool) {
	return getImageUrl(configuration.Images.ProfileSizes, configuration.Images.BaseUrl, size, string(p))
}

func (p ProfileImage) GetSecureUrl(configuration *Configuration, size ProfileSize) (string, bool) {
	return getImageUrl(configuration.Images.ProfileSizes, configuration.Images.SecureBaseUrl, size, string(p))
}

type StillImage string

func (s StillImage) GetUrl(configuration *Configuration, size StillSize) (string, bool) {
	return getImageUrl(configuration.Images.StillSizes, configuration.Images.BaseUrl, size, string(s))
}

func (s StillImage) GetSecureUrl(configuration *Configuration, size StillSize) (string, bool) {
	return getImageUrl(configuration.Images.StillSizes, configuration.Images.SecureBaseUrl, size, string(s))
}

func getImageUrl[SizeType ~string, SizeListType []SizeType](sizeList SizeListType, base string, size SizeType, rawPath string) (string, bool) {
	for _, s := range sizeList {
		if s == size {
			return base + string(size) + rawPath, true
		}
	}
	return "", false
}