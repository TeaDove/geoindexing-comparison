import streamlit as st
import pandas as pd

points = pd.read_csv("points.csv")

st.map(points)
