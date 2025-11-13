# serpa-cli

CLI tool to upload categories, places and assets to the serpa maps backend

## usage:

`categories.csv` and `places.csv` must exist. In order to add assets to a place they should be in a folder with the exact place name. The order in which categories and places are added is defined by their position in the csv. The first column is added first. Same applies for assets inside a folder.

```
folder/
├── categories.csv
├── places.csv
└── Null Island
    ├── first_image.jpg
    └── second_image.png
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
