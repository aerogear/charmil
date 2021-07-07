# Charmil
<a href="https://pkg.go.dev/github.com/aerogear/charmil"><img src="https://pkg.go.dev/badge/github.com/aerogear/charmil.svg" alt="Go Reference"></a>
![example workflow](https://github.com/aerogear/charmil/actions/workflows/go_check.yaml/badge.svg)
<a href="https://discord.gg/nAQBYZncvm">
  <img alt="Discord Server" src="https://img.shields.io/discord/632220458137419776?logo=Discord&logoColor=%23fff">
</a>


<p align="center">
  <img width="400" src="https://github.com/aerogear/charmil/raw/main/resources/logo.png">
  <br/>
  Framework for building command line plugins on top of Golang Cobra Library.  <br/>
  Charmil will let you control your cobra based CLIs. 🚀
</p>


## Introduction

Charmil provides an ecosystem to build production ready command line tools using the Cobra Framework. It is a SDK which
allows for the creation of CLIs that all adhere to a similarly structured architecture.
It provides a level of elasticity to add and/ or remove a plugin CLI or subset of commands from a plugin CLI to a main 'host' CLI.


### Architecture

Charmil provides a way to build a CLI ecosystem where developers can build multiple fragmented CLI's in various repositories 
and embed them later in the single host CLI. Charmil provides core base functionality that helps to abstract various elements 
of the CLI ecosystem:

- Logging
- Authentication
- Configuration
- Commands and Flags standards.

Charmil SDK will typically be introduced into two different CLIs:

- Host CLI - CLI that is being used by end user that can embed one or more plugin CLIs
- Plugin CLIs - Separate CLIs that can be embedded into the host CLI


## Components
<p align="center">
  <img width="400" src="https://github.com/aerogear/charmil/raw/main/resources/charmil-base-ssv-3-pillar.png">
</p>

### Charmil Plugin framework starter template 

Creating Golang CLI should be easy and fast. 
Charmil starter template provides fast way to build plugin and host commands that use Charmil SDK and Validator.
Starter template contains all tools and solutions required to build command line tool

### Charmil SDK

Framework for building modular, multi repo Golang CLI's with Cobra 
Number of high level implementations for individual CLI creators for common problems like configuration, authentication and internationalization
Provides ability to embed modular CLI into hosts CLI at compile time.
 
### Charmil Validator

Charmil Validator gives developers the ability to validate a group of their cobra commands against common practices. 
Validator can check if commands are providing proper documentation, doesn't have overriden flags or provide shell completions. 
Validator can be used as go unit test and run during CI/CD that simplifies review of CLI's that contain commands that are hosted in multiple repositories.

 
### Charmil Command Registry

Charmil command registry lets you provide external index of the command line tools that can be installed dynamically and embeeded into your host CLI.
Developers can create CLI profiles to install multiple individual CLI's at the same time and keep them updated with various backends like Kubernetes/OpenShift etc.

> Under Construction

## Documentation

Visit the [Charmil Documentation](https://aerogear.github.io/charmil/docs/)

## License

Licensed under the Apache License 2.0
