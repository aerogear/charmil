"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[322],{3905:function(e,t,n){n.d(t,{Zo:function(){return c},kt:function(){return p}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),u=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=u(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),m=u(n),p=a,h=m["".concat(s,".").concat(p)]||m[p]||d[p]||i;return n?r.createElement(h,o(o({ref:t},c),{},{components:n})):r.createElement(h,o({ref:t},c))}));function p(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=m;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var u=2;u<i;u++)o[u]=n[u];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},2434:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return s},metadata:function(){return u},toc:function(){return c},default:function(){return m}});var r=n(7462),a=n(3366),i=(n(7294),n(3905)),o=["components"],l={title:"Charmil Validator",slug:"/charmil_validator"},s=void 0,u={unversionedId:"charmil_validator",id:"charmil_validator",isDocsHomePage:!1,title:"Charmil Validator",description:"Charmil Validator can be used for testing and controlling many aspects of cobra commands. It provides many rules out of the box for validating the commands.",source:"@site/../docs/src/charmil_validator.md",sourceDirName:".",slug:"/charmil_validator",permalink:"/charmil/docs/charmil_validator",version:"current",frontMatter:{title:"Charmil Validator",slug:"/charmil_validator"},sidebar:"main",previous:{title:"Charmil Example Builder",permalink:"/charmil/docs/charmil_example_builder"}},c=[{value:"Rules provided by validator",id:"rules-provided-by-validator",children:[]},{value:"How to use",id:"how-to-use",children:[]},{value:"Disable a Rule",id:"disable-a-rule",children:[]},{value:"Ignore Commands",id:"ignore-commands",children:[{value:"Skip single command <code>mycli actions create</code>",id:"skip-single-command-mycli-actions-create",children:[]},{value:"Skip the command including all children",id:"skip-the-command-including-all-children",children:[]},{value:"Skip the command for specific rule",id:"skip-the-command-for-specific-rule",children:[]}]}],d={toc:c};function m(e){var t=e.components,l=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},d,l,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"Charmil Validator can be used for testing and controlling many aspects of cobra commands. It provides many rules out of the box for validating the commands."),(0,i.kt)("p",null,(0,i.kt)("img",{alt:"charmil validator",src:n(4765).Z})),(0,i.kt)("h2",{id:"rules-provided-by-validator"},"Rules provided by validator"),(0,i.kt)("h4",{id:"lengthrule"},"LengthRule"),(0,i.kt)("p",null,"Length Rule can control the lengths of strings in different attributes of cobra.Command structure."),(0,i.kt)("h4",{id:"mustexistrule"},"MustExistRule"),(0,i.kt)("p",null,"Must Exist Rule ensures that the selected attributes from ",(0,i.kt)("inlineCode",{parentName:"p"},"cobra.Command")," strcuture must be present in the command."),(0,i.kt)("h4",{id:"usematchesrule"},"UseMatchesRule"),(0,i.kt)("p",null,"Use Matches Rule compares the Use attribute of ",(0,i.kt)("inlineCode",{parentName:"p"},"cobra.Command")," with the user provided regexp."),(0,i.kt)("h4",{id:"examplematchesrule"},"ExampleMatchesRule"),(0,i.kt)("p",null,"Example Matches Rule ensures that the command is properly documented with the proper examples in it."),(0,i.kt)("h4",{id:"punctuationrule"},"PunctuationRule"),(0,i.kt)("p",null,"Punctuation Rule checks for the punctuation errors in the command according to industry standards."),(0,i.kt)("h2",{id:"how-to-use"},"How to use"),(0,i.kt)("p",null,"It is recommended to use the validator while writing unit tests for cobra commands."),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Create a configuration of type ",(0,i.kt)("inlineCode",{parentName:"li"},"rules.ValidatorConfig"),". You can provide your own ValidatorConfig or use the default one by leaving it empty")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/aerogear/charmil/validator"\n\nvar ruleCfg rules.ValidatorConfig\n')),(0,i.kt)("p",null,"or overriding default config"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'import "github.com/aerogear/charmil/validator"\n\nruleCfg := rules.ValidatorConfig{\n    ValidatorRules: rules.ValidatorRules{\n        Length: rules.Length{Limits: map[string]rules.Limit{"Use": {Min: 1}}},\n        MustExist: rules.MustExist{Fields: map[string]bool{"Args": true}},\n        UseMatches: rules.UseMatches{Regexp: `^[^-_+]+$`},\n    },\n}\n')),(0,i.kt)("ol",{start:2},(0,i.kt)("li",{parentName:"ol"},"Generate the validation errors by using ",(0,i.kt)("inlineCode",{parentName:"li"},"ExecuteRules")," function over the ruleCfg")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"validationErr := rules.ExecuteRules(cmd, &ruleCfg)\n")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"ExecuteRules")," function will return a slice of ",(0,i.kt)("inlineCode",{parentName:"p"},"ValidationError")," object, which can be efficiently used for testing purposes."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'if len(validationErr) != 0 {\n    t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)\n}\nfor _, errs := range validationErr {\n    if errs.Err != nil {\n        t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)\n    }\n}\n')),(0,i.kt)("h2",{id:"disable-a-rule"},"Disable a Rule"),(0,i.kt)("p",null,"All the rules provided by charmil are enabled by default. If you want to turn off particular rule or rules, there is an ",(0,i.kt)("inlineCode",{parentName:"p"},"Disable")," option in ",(0,i.kt)("inlineCode",{parentName:"p"},"RuleOptions")," in each rule."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"ruleCfg := rules.ValidatorConfig{\n    ValidatorRules: rules.ValidatorRules{\n        Punctuation: Punctuation{\n            RuleOptions: validator.RuleOptions{\n                Disable: true,\n            },\n        },\n        UseMatches: UseMatches{\n            RuleOptions: validator.RuleOptions{\n                Disable: true,\n            },\n        },\n\n    },\n}\n")),(0,i.kt)("h2",{id:"ignore-commands"},"Ignore Commands"),(0,i.kt)("p",null,"Sometimes during development, you want to pass the tests for certain commands, but at the same time use Validator for tests. Validation can be skipped/ignored for the commands, mentioned in the validator configuration.\nTo ignore the commands you need to specify the path of the command in validator configuration."),(0,i.kt)("p",null,"Assume your CLI is having a command structure like this:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"mycli\n    actions\n        create\n        update\n        delete\n        read\n")),(0,i.kt)("h3",{id:"skip-single-command-mycli-actions-create"},"Skip single command ",(0,i.kt)("inlineCode",{parentName:"h3"},"mycli actions create")),(0,i.kt)("p",null,"Use ",(0,i.kt)("inlineCode",{parentName:"p"},"SkipCommands")," option in ",(0,i.kt)("inlineCode",{parentName:"p"},"ValidatorOptions")," to skip validation for a command"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'ValidatorOptions: rules.ValidatorOptions{\n    SkipCommands: map[string]bool{"mycli actions create": true},\n},\n')),(0,i.kt)("h3",{id:"skip-the-command-including-all-children"},"Skip the command including all children"),(0,i.kt)("p",null,"Use ",(0,i.kt)("inlineCode",{parentName:"p"},"SkipCommands")," option in ",(0,i.kt)("inlineCode",{parentName:"p"},"ValidatorOptions")," using asterisk sign to skip validation for all subcommands"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'ValidatorOptions: rules.ValidatorOptions{\n    SkipCommands: map[string]bool{"mycli actions*": true},\n},\n')),(0,i.kt)("h3",{id:"skip-the-command-for-specific-rule"},"Skip the command for specific rule"),(0,i.kt)("p",null,"Use ",(0,i.kt)("inlineCode",{parentName:"p"},"SkipCommands")," option in ",(0,i.kt)("inlineCode",{parentName:"p"},"RuleOptions")," to skip validation for specific rule"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'Length: rules.Length{\n    RuleOptions: validator.RuleOptions{\n        SkipCommands: map[string]bool{"mycli actions create": true},\n    },\n    Limits: map[string]rules.Limit{\n        "Use": {Min: 1},\n    },\n},\n')))}m.isMDXComponent=!0},4765:function(e,t,n){t.Z=n.p+"assets/images/charmil_validator-95ca0451dcd8e691e9a1a49133ff0e04.png"}}]);