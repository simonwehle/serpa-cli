# serpa-cli

CLI tool to upload categories, places and assets to the serpa maps backend

## usage:

The following folder structure is required

```
folder/
├── categories.csv
├── places.csv
└── Test
    └── test_asset.jpg
```

`categories.csv` should contain the following columns

```csv
name,icon,color
landmark,camera_alt,#9c27b0"
```

`places.csv` should contain the following columns (description can be empty)

```csv
latitude,longitude,name,description,category
0,0,Null Island,,landmark
```

## build

```sh
go build .
sudo mv serpa-cli /usr/local/bin/
```

uninstall

```
sudo rm /usr/local/bin/serpa-cli
```
