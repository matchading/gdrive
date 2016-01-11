package cli

// TODO
// Default values? Default string values? Parser must always return a value
// Support invalid flag combinations?


type Flag interface {
    GetPatterns() []string
    GetName() string
    GetDescription() string
    GetParser() Parser
}

func getFlagParser(flags []Flag) Parser {
    var parsers []Parser

    for _, flag := range flags {
        parsers = append(parsers, flag.GetParser())
    }

    return FlagParser{parsers}
}


type BoolFlag struct {
    Patterns []string
    Name string
    Description string
    DefaultValue bool
    OmitValue bool
}

func (self BoolFlag) GetName() string {
    return self.Name    
}

func (self BoolFlag) GetPatterns() []string {
    return self.Patterns
}

func (self BoolFlag) GetDescription() string {
    return self.Description
}

func (self BoolFlag) GetParser() Parser {
    var parsers []Parser
    for _, p := range self.Patterns {
        parsers = append(parsers, BoolFlagParser{
            pattern: p,
            key: self.Name,
            omitValue: self.OmitValue,
            defaultValue: self.DefaultValue,
        })
    }

    if len(parsers) == 1 {
        return parsers[0]
    }
    return ShortCircuitParser{parsers}
}


type StringFlag struct {
    Patterns []string
    Name string
    Description string
    DefaultValue string
}

func (self StringFlag) GetName() string {
    return self.Name    
}

func (self StringFlag) GetPatterns() []string {
    return self.Patterns
}

func (self StringFlag) GetDescription() string {
    return self.Description
}

func (self StringFlag) GetParser() Parser {
    var parsers []Parser
    for _, p := range self.Patterns {
        parsers = append(parsers, StringFlagParser{
            pattern: p,
            key: self.Name,
            defaultValue: self.DefaultValue,
        })
    }

    if len(parsers) == 1 {
        return parsers[0]
    }
    return ShortCircuitParser{parsers}
}

type IntFlag struct {
    Patterns []string
    Name string
    Description string
    DefaultValue int64
}

func (self IntFlag) GetName() string {
    return self.Name    
}

func (self IntFlag) GetPatterns() []string {
    return self.Patterns
}

func (self IntFlag) GetDescription() string {
    return self.Description
}

func (self IntFlag) GetParser() Parser {
    var parsers []Parser
    for _, p := range self.Patterns {
        parsers = append(parsers, IntFlagParser{
            pattern: p,
            key: self.Name,
            defaultValue: self.DefaultValue,
        })
    }

    if len(parsers) == 1 {
        return parsers[0]
    }
    return ShortCircuitParser{parsers}
}
