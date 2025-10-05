# Chase's Simple Archiver

After one too many times borking my operating system, I decided to build a basic tool to back up the important stuff on my machine. I'm positive there is some much better tools out there that do the same thing but I felt like it was easier to write one myself than to try and understand (and trust) something new.

## No BS Explanation

This is just some Go code to upload directories and files you put in a list to an S3 bucket.

## Usage:

`csa init` - Create initial `archive.yml` (explained in next section).

`csa sync` - Syncs local archive with remote archive.

`csa restore` - Pulls latest remote archive and attempts to load it into target paths. Will ask you to confirm before overwriting anything.

`csa add /path/to/target` - Add target to list of directories/files that will be archived during a sync.

`csa remove /path/to/target` - Remove target from list of directories/files that will be archived during a sync.

`csa exclude /path/to/target <pattern>` - Add exclusion (glob) pattern to a target.

`csa list` - List all currently tracked targets.

`csa remote s3:<bucket>` - Initialize/configure S3 remote.

## Configuration

`~/.config/csa/archive.yml` - determines what directories/files are archived during a sync, along with any other relevant metadata. This repo also contains an [`example.archive.yml`](example.archive.yml) if you're curious about the format.

## Contributing

I don't expect anyone to use this other than me but if for some reason you do and want to contribute, PRs are open.

## Roadmap

- Symlink support (currently just skips them)
- Automatic sync
- Non s3 remotes
- Customizable config location
