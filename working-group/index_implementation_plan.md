
# Charmil Index Implementation Plan

- ## Possible methods to reuse Krew components

	1. **Directly use as library** (Cannot be followed when the required component is unexportable or is present in the `internal` directory).

	2. **Copy and use the required component after making only some trivial modifications** (while following the Apache 2.0 license guidelines).
	
	3. **Copy and use the required component after making some major, necessary modifications according to the use case** (while following the Apache 2.0 license guidelines).

  &nbsp;
    However, for smaller components, we can consider using methods 2 or 3 instead of method 1.

   > **_A little copying is better than a little dependency._**
   >
   > &mdash; <cite>Rob Pike [(Gopherfest'15)](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s)</cite>
	
  &nbsp;
 - ## Krew-index Features

	***The checked features are the ones that we can include in our MVP.***

	- [x] Plugins can be packaged as `.zip` or `.tar.gz` archives
	- [x] Add a new plugin by tagging a git release of the plugin archive with a [semantic version](https://semver.org/) (e.g. `v1.0.0`) and then creating a PR to the index repository.
	- [ ] Uses [krew-index-autoapprove](https://github.com/ahmetb/krew-index-autoapprove) bot to automatically merge PRs involving plugin updates in the index repository. 
	Example: [here](https://github.com/kubernetes-sigs/krew-index/pull/508)
	- [ ] Completely automated plugin update process, achieved by integrating the [krew-index-autoapprove](https://github.com/ahmetb/krew-index-autoapprove) bot with the [krew-release-bot](https://github.com/rajatjindal/krew-release-bot), which is Github Action that automatically bumps the version in  `krew-index`  repo every time someone pushes a new git tag to the plugin repository and creates a PR on our behalf.
	Example: [here](https://github.com/kubernetes-sigs/krew-index/pull/490)

	- [x] Option to use Custom Plugin Indexes
	- [ ] Shows plugin usage analytics on the [stats.krew.dev](https://datastudio.google.com/c/reporting/f74370a0-adcf-4cec-b7bd-a58c638948f5/page/Ufl7) dashboard, obtained by scraping the downloads count of plugin assets via the [GitHub API](https://developer.github.com/v3/repos/releases/#list-assets-for-a-release) regularly.

&nbsp;
 - ## Krew Reusable Code Components


	|   Include in the MVP?	|   Component Name 	|   Description	|   Method to use (Refer to first bullet point)	|
	|---	|---	|---	|---	|
	|   Yes	|  [validate-krew-manifest](https://github.com/kubernetes-sigs/krew/tree/master/cmd/validate-krew-manifest) 	|   Makes sure a manifest file is valid	|   2	|
	|   Yes	|   [downloader.go](https://github.com/kubernetes-sigs/krew/blob/master/internal/download/downloader.go)	|   Contains methods that are responsible for pulling the plugin URI, verifying it, downloading and then extracting the `.zip` or `.tar.gz` archive to the destination folder for installation	|   2	|
	|   No	|   [krew-index-autoapprove](https://github.com/ahmetb/krew-index-autoapprove) bot	|   ^ Mentioned in the features section above	|   3	|
	|   No	|   [krew-release-bot](https://github.com/rajatjindal/krew-release-bot)	|   ^ Mentioned in the features section above	|   3	|
	|   No	|   [krew-index-tracker](https://github.com/corneliusweig/krew-index-tracker)	|   A scraper tool which tracks the download counts of `krew` plugins using GitHub API	|   3	|


	PS: Since this file is limited to discussion on Index, only the code components related to the Krew-index features have been mentioned here.
&nbsp;
 - ## Some important points to discuss
	 - **Whether we want our index working the same as that of Krew for SDK**
		 - Due to some variations in the Charmil and Krew models (as discussed in [charmil_personas.md](https://github.com/aerogear/charmil/blob/main/working-group/charmil_personas.md)), the index in Charmil SDK will work a bit differently.
		 - In contrast to Krew, there will be no such thing as a default index here.
		 - It is the developers who will be given the option to add their custom indexes here, unlike Krew where end-users do the same.
		 - In the MVP, we can allow the developers to hardcode their indexes. 
		 - At a later stage, we can add the feature through which the developers can manage (list, add, update, delete) custom indexes through a CLI (just the way it is done for end-users in Krew).
		 
	 - **Should we support packaged files or single binary?**
		 - Currently, Krew supports only packaged archive files as plugins. 
		 - As mentioned above, as far as our MVP is concerned, we can reuse the [downloader.go](https://github.com/kubernetes-sigs/krew/blob/master/internal/download/downloader.go) component of Krew which will take care of the installation process for us.
		 - At a later stage, we can consider adding support for single binary files too.