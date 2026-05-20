#import "@preview/js:0.1.3": *
#show: js.with()
#show math.equation.where(block: true): set block(inset: (left: 2em))
#show math.equation.where(block: true): set align(left)
#show "、": "，"
#show "。": "．"

#maketitle(
  title: "レポート課題 1",
  authors: "1027373499 野田悠晟"
)

= 血液型に関する分析
== 各血液型の人数および割合
#table(
  align: center,
  columns: 3,

  [血液型],[人数 (人)],[割合 (%: 四捨五入)],
  [A], [65], [34.8],
  [B], [37], [19.8],
  [O], [57], [6.4],
  [AB], [12], [30.5],
  [others], [16], [8.6],
  [計], [187], [100]
)

== グラフによる可視化
=== 棒グラフ
#image("images/bloodtype_bar.png")
\
\
=== 帯グラフ
#image("images/bloodtype_band.png")
\
\
=== 円グラフ
#image("images/bloodtype_pie.png")
#pagebreak()

= お菓子の好みに関する分析
== 各調査項目におけるグラフ可視化
=== 調査項目 1 : "きのこの山とたけのこの里，どちらが好きですか?"

棒グラフ\
#image("images/kinokoOrTakenoko_bar.png")
帯グラフ\
#image("images/kinokoOrTakenoko_band.png")
円グラフ\
#image("images/kinokoOrTakenoko_pie.png")

#pagebreak()

=== 調査項目 2 : "きのこの山とたけのこの里，両方好きですか?"

棒グラフ\
#image("images/bothOrNot_bar.png")
帯グラフ\
#image("images/bothOrNot_band.png")
円グラフ\
#image("images/bothOrNot_pie.png")

#pagebreak()

=== 調査項目 3 : "チョコレートは好きですか? (1～5 の 5 段階評価)"

棒グラフ\
#image("images/chocolate_bar.png")
帯グラフ\
#image("images/chocolate_band.png")
円グラフ\
#image("images/chocolate_pie.png")

#pagebreak()

== 調査項目 1 および 2 におけるクロス集計表
#table(
  align: center,
  columns: 4,

  [], [両方いける], [片方だけ], [計],
  [きのこの山が好き], [61], [17], [78],
  [たけのこの里が好き], [93], [15], [108],
  [計], [154], [32], [186]
  )
\
なお、この結果から導き出せる推定値は以下の表のようになる。\
\
#table(
  align: center,
  columns: 4,

  [], [両方いける], [片方だけ], [計],
  [きのこの山が好き], [64.6], [13.4], [78],
  [たけのこの里が好き], [89.4], [18.6], [108],
  [計], [154], [32], [186]
)

#pagebreak()

= 2022年における京都の気温に関する分析
== 各日の平均・最高・および最低気温の統計的解析
=== 平均気温について (℃)
\
#table(
  align: center,
  columns: 8,

  [平均], [分散], [標準偏差], [最小値], [第一四分位数], [中央値], [第三四分位数], [最大値],
  [16.9], [81.6], [9.0], [1.4], [8.5], [16.8], [25.5], [32]
)
\
=== 最高気温について (℃)
\
#table(
  align: center,
  columns: 8,

  [平均], [分散], [標準偏差], [最小値], [第一四分位数], [中央値], [第三四分位数], [最大値],
  [21.9], [88.7], [9.4], [3.7], [13.1], [22.2], [30.2], [38.6]
)
\
=== 最低気温について (℃)
\
#table(
  align: center,
  columns: 8,

  [平均], [分散], [標準偏差], [最小値], [第一四分位数], [中央値], [第三四分位数], [最大値],
  [12.8], [81.7], [9.0], [-1.8], [4.2], [12.6], [22.2], [27.9]
)

#pagebreak()

== ヒストグラムによる可視化
\
いずれも、縦軸に割合、横軸に気温(℃)をとっている。\
\
=== 平均気温
#image("images/temp_average.png")

=== 最高気温
#image("images/temp_max.png")

=== 最低気温
#image("images/temp_min.png")

#pagebreak()

== 折れ線グラフによる可視化
\
赤は各日の平均値、緑は最高気温、青は最低気温を示している。
#image("images/temperature_line.png")

#pagebreak()

== 最高気温と最低気温の差の可視化
=== ヒストグラム
縦軸に割合、横軸に気温(℃)をとっている。\
\
#image("images/temp_diff.png")

=== 折れ線グラフ
#image("images/temperature_diff_line.png")
