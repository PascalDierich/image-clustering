# Image-Clustering

The _image-clustering_ program partitions the pixels in an image
based on their color to a pre-defined number of clusters.

```
Usage:
  -in string
        Path to input image
  -k int
        Number of cluster (default 10)
  -out string
        Path to save clustered image (default "${HOME}/img_clustered.png")
```
### Examples:

#### original
![img4](data/img4.png)

#### k = 5
![img4_c5](data/img4_c05.png)

#### k = 100
![img4_c100](data/img4_c100.png)

#### k = 500
![img4_c500](data/img4_c500.png)
