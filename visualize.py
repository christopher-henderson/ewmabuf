#!/usr/bin/env python
import json

import plotly.graph_objects as go

# Create random data with numpy
import numpy as np



def normal_distribution(size, loc, scale) -> [int]:
    # size=1000, loc=1024*1024*2, scale=1024*1024
    from numpy import random
    from pstream import Stream
    from math import floor
    return Stream(random.normal(size=size, loc=loc, scale=scale)).map(lambda x: max(x, 0)).map(lambda x: min(x, 1024*1024*4)).map(floor).collect()


def load():
    path = "/home/chris/projects/ewmabuf/normal_result.json"
    with open(path, "r") as f:
        results = json.load(f)
    print(results)

    x = [i + 1 for i in range(len(results['Messages']))]

    fig = go.Figure()
    fig.add_trace(go.Scatter(x=x, y=results['Messages'],
                             mode='lines',
                             name='messages'))
    fig.add_trace(go.Scatter(x=x, y=results['BufSize'],
                             mode='lines',
                             name='buf size'))
    fig.add_trace(go.Scatter(x=x, y=results['AverageOverAlloc'],
                             mode='lines', name='average overalloc'))

    fig.show()

load()

def sample():
    np.random.seed(1)

    N = 100
    random_x = np.linspace(0, 1, N)
    random_y0 = np.random.randn(N) + 5
    random_y1 = np.random.randn(N)
    random_y2 = np.random.randn(N) - 5

    # Create traces
    fig = go.Figure()
    fig.add_trace(go.Scatter(x=random_x, y=random_y0,
                             mode='lines',
                             name='lines'))
    fig.add_trace(go.Scatter(x=random_x, y=random_y1,
                             mode='lines+markers',
                             name='lines+markers'))
    fig.add_trace(go.Scatter(x=random_x, y=random_y2,
                             mode='markers', name='markers'))

    fig.show()
