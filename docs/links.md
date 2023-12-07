Материалы:
- [Spatial Indexes](https://db.in.tum.de/downloads/publications/learnedspatial.pdf)
- Индексы:
	- Tree:
		- R-based:
			- R Tree [wiki](https://en.wikipedia.org/wiki/R-tree), [goimpl](https://github.com/tidwall/rtree), [goimpl_2](https://github.com/dhconnelly/rtreego)
			- R Star Tree [wiki](https://en.wikipedia.org/wiki/R*-tree), https://habr.com/ru/articles/666904/
			- Hubert R Tree
		- Quad tree [wiki](https://en.wikipedia.org/wiki/Quadtree), [goimpl](https://github.com/JamesLMilner/quadtree-go)
		- K-d tree [wiki](https://en.wikipedia.org/wiki/K-d_tree), [goimpl](https://github.com/kyroy/kdtree)
        - BSP tree https://ru.wikipedia.org/wiki/%D0%94%D0%B2%D0%BE%D0%B8%D1%87%D0%BD%D0%BE%D0%B5_%D1%80%D0%B0%D0%B7%D0%B1%D0%B8%D0%B5%D0%BD%D0%B8%D0%B5_%D0%BF%D1%80%D0%BE%D1%81%D1%82%D1%80%D0%B0%D0%BD%D1%81%D1%82%D0%B2%D0%B0
        - VP tree https://ru.wikipedia.org/wiki/VP-%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE
	- Geohash:
		- Geohash [wiki](https://en.wikipedia.org/wiki/Geohash), [goimpl](https://github.com/mmcloughlin/geohash)
		- S2, [docs](https://s2geometry.io/), [goimpl](https://pkg.go.dev/github.com/golang/geo/s2)
		- H3, [docs](https://www.uber.com/blog/h3/), [goimpl](https://github.com/uber/h3-go)
    - Range:
  		- BRIN
- TeX: [overleaf](https://www.overleaf.com/project/64594cfe9c8fa3c587c5d604)
- Шизоиндексы от корейцев
  -  The proposed
	 indexes typically took an R-tree or B-tree-like structure, such as P-tree [ 26 ], Trajtree [27 ], Trails-tree [28 ],
	 DITIR [29], V-tree [30], and D-Toss [31]

Литература:
- https://onlinelibrary.wiley.com/doi/10.1002/cpe.6029
- https://sci-hub.ru/https://onlinelibrary.wiley.com/doi/10.1002/cpe.6029
- https://www.mdpi.com/2071-1050/12/22/9727

    A. Guttman. R-trees: A Dynamic Index Structure for Spatial Searching. Proceedings of ACM SIGMOD, pages 47-57, 1984. http://www.cs.jhu.edu/~misha/ReadingSeminar/Papers/Guttman84.pdf

    N. Beckmann, H .P. Kriegel, R. Schneider and B. Seeger. The R*-tree: An Efficient and Robust Access Method for Points and Rectangles. Proceedings of ACM SIGMOD, pages 323-331, May 1990. http://infolab.usc.edu/csci587/Fall2011/papers/p322-beckmann.pdf

    N. Roussopoulos, S. Kelley and F. Vincent. Nearest Neighbor Queries. ACM SIGMOD, pages 71-79, 1995. http://www.postgis.org/support/nearestneighbor.pdf

- dimploma:
- https://drive.google.com/drive/folders/1EuwImS3tjUvugrCW3wSV1M3QoD8rBBe8?usp=drive_link
