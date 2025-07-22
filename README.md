# snaglr
File ingestion tool for hoarders and archivers

CLI structure

```txt
ingestr
├── --help
├── --version
├── --log-level string              # info (default), debug, error
├── --quiet                         # suppress console output, log to file only
├── init                            # generate sample config.yaml
├── cleanup                         # remove files from managed temporary locations
└── run
    ├── --config, -c string         # path to config file (default ~/.config/ingestr/config.toml)
    ├── --input, -i string[]        # one or more input folders to scan
    ├── --recurse, -r               # recurse into folders
    ├── --allow-ext, -a string[]    # allowed file extensions (e.g. jpg,png)
    ├── --max-size, -s int          # maximum file size in MB
    ├── --preset, -p string         # export preset name from config
    ├── --move                      # move instead of copy
    ├── --dest, -d string           # destination folder
    ├── --force, -f                 # force overwrite if file exists
    └── --dry-run                   # preview only, do not write
```