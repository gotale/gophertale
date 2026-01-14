package protocol

import "image/color"

type Weather struct {
	ID                        string
	TaxIndexes                []int32
	Stars                     string
	Moons                     map[int32]string
	Clouds                    []Cloud
	SunlightDampingMultiplier map[float32]float32
	SunlightColours           map[float32]color.RGBA
	SkyTopColours             map[float32]color.RGBA
	SkyBottomColours          map[float32]color.RGBA
	SkySunsetColours          map[float32]color.RGBA
	SunColours                map[float32]color.RGBA
	SunScales                 map[float32]float32
	SunGlowColours            map[float32]color.RGBA
	MoonColours               map[float32]color.RGBA
	MoonScales                map[float32]float32
	MoonGlowColours           map[float32]color.RGBA
	FogColours                map[float32]color.RGBA
	FogHeightFalloffs         map[float32]float32
	FogDensities              map[float32]float32
	ScreenEffect              string
	ScreenEffectColours       map[float32]color.RGBA
	ColourFilters             map[float32]color.RGBA
	WaterTints                map[float32]color.RGBA
	Particle                  WeatherParticle
	Fog                       NearFar
	FogOptions                FogOptions
}

func (x *Weather) Marshal(io IO) {
	opts := NewOptionals(io)
	hasId := opts.Has(x.ID != "")
	hasTagIndexes := opts.Has(len(x.TaxIndexes) > 0)
	hasStars := opts.Has(x.Stars != "")
	hasMoons := opts.Has(len(x.Moons) > 0)
	hasClouds := opts.Has(len(x.Clouds) > 0)
	hasSunlightDamping := opts.Has(len(x.SunlightDampingMultiplier) > 0)
	hasSunlightColours := opts.Has(len(x.SunlightColours) > 0)
	hasSkyTopColours := opts.Has(len(x.SkyTopColours) > 0)
	opts.Write()

	opts = NewOptionals(io)
	hasSkyBottomColours := opts.Has(len(x.SkyBottomColours) > 0)
	hasSkySunsetColours := opts.Has(len(x.SkySunsetColours) > 0)
	hasSunColours := opts.Has(len(x.SunColours) > 0)
	hasSunScales := opts.Has(len(x.SunScales) > 0)
	hasSunGlowColours := opts.Has(len(x.SunGlowColours) > 0)
	hasMoonColours := opts.Has(len(x.MoonColours) > 0)
	hasMoonScales := opts.Has(len(x.MoonScales) > 0)
	hasMoonGlowColours := opts.Has(len(x.MoonGlowColours) > 0)
	opts.Write()

	opts = NewOptionals(io)
	hasFogColours := opts.Has(len(x.FogColours) > 0)
	hasFogHeightFalloffs := opts.Has(len(x.FogHeightFalloffs) > 0)
	hasFogDensities := opts.Has(len(x.FogDensities) > 0)
	hasScreenEffect := opts.Has(x.ScreenEffect != "")
	hasScreenEffectColours := opts.Has(len(x.ScreenEffectColours) > 0)
	hasColourFilters := opts.Has(len(x.ColourFilters) > 0)
	hasWaterTints := opts.Has(len(x.WaterTints) > 0)
	hasParticle := opts.Has(x.Particle != WeatherParticle{})
	opts.Write()

	opts = NewOptionals(io)
	hasFog := opts.Has(x.Fog != NearFar{})
	hasFogOptions := opts.Has(x.FogOptions != FogOptions{})
	opts.Write()

	// TODO: Encoding
	_ = hasId
	_ = hasTagIndexes
	_ = hasStars
	_ = hasMoons
	_ = hasClouds
	_ = hasSunlightDamping
	_ = hasSunlightColours
	_ = hasSkyTopColours
	_ = hasSkyBottomColours
	_ = hasSkySunsetColours
	_ = hasSunColours
	_ = hasSunScales
	_ = hasSunGlowColours
	_ = hasMoonColours
	_ = hasMoonScales
	_ = hasMoonGlowColours
	_ = hasFogColours
	_ = hasFogHeightFalloffs
	_ = hasFogDensities
	_ = hasScreenEffect
	_ = hasScreenEffectColours
	_ = hasColourFilters
	_ = hasWaterTints
	_ = hasParticle
	_ = hasFog
	_ = hasFogOptions
}

type Cloud struct {
	Texture string
	Speeds  map[float32]float32
	Colours map[float32]color.RGBA
}

func (x *Cloud) Marshal(io IO) {
	opts := NewOptionals(io)
	hasTexture := opts.Has(x.Texture != "")
	hasSpeeds := opts.Has(len(x.Speeds) > 0)
	hasColours := opts.Has(len(x.Colours) > 0)
	opts.Write()

	// TODO: Encoding
	_ = hasTexture
	_ = hasSpeeds
	_ = hasColours
}

type WeatherParticle struct {
	SystemID                 string
	Colour                   color.RGBA
	Scale                    float32
	OvergroundOnly           bool
	PositionOffsetMultiplier float32
}

func (x *WeatherParticle) Marshal(io IO) {
	opts := NewOptionals(io)
	hasSystemID := opts.Has(x.SystemID != "")
	opts.Has(x.Colour != color.RGBA{})
	opts.Write()

	io.RGB(&x.Colour)
	io.Float32(&x.Scale)
	io.Bool(&x.OvergroundOnly)
	io.Float32(&x.PositionOffsetMultiplier)
	if hasSystemID {
		io.VarString(&x.SystemID, 4096000)
	}
}

type NearFar struct {
	Near float32
	Far  float32
}

func (x *NearFar) Marshal(io IO) {
	io.Float32(&x.Near)
	io.Float32(&x.Far)
}

type FogOptions struct {
	IgnoreFogLimits                 bool
	EffectiveViewDistanceMultiplier float32
	FogFarViewDistance              float32
	FogHeightCameraOffset           float32
	FogHeightCameraOverridden       bool
	FogHeightCameraFixed            float32
}

func (x *FogOptions) Marshal(io IO) {
	io.Bool(&x.IgnoreFogLimits)
	io.Float32(&x.EffectiveViewDistanceMultiplier)
	io.Float32(&x.FogFarViewDistance)
	io.Float32(&x.FogHeightCameraOffset)
	io.Bool(&x.FogHeightCameraOverridden)
	io.Float32(&x.FogHeightCameraFixed)
}
