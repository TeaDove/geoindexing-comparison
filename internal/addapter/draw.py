#!/usr/bin/python3

import pandas as pd
import matplotlib.pyplot as plt
import sys


if __name__ == "__main__":
    df = pd.read_csv(sys.argv[1])
    ax = plt.gca()

    for color in df["color"].unique():
        colored_df = df[df["color"] == color]
        colored_df.plot(
            kind="scatter",
            x="lon",
            y="lat",
            grid=True,
            stacked=True,
            ax=ax,
            color=color,
            s=10,
        )
    plt.savefig(sys.argv[2])
