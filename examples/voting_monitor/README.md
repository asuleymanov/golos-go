# Voting Monitor

In this example we connect to `golosd` and watch operations as they
are happening. Every time we see a `vote` operation, we print a message
into the console.

```
$ go run main.go 
2018/03/05 12:55:17 ---> GetConfig()
2018/03/05 12:55:17 ---> Entering the block processing loop (last block = 14429786)
2018/03/05 12:55:20 @freelancer1 voted for @refreshka/berech-klimat-i-zarabatyvat-eto-vozmozhno-s-earth-token
2018/03/05 12:55:20 @flinth1 voted for @golos.loto/5x36-golos-lottery-1622-8106
2018/03/05 12:55:20 @freelancer1 voted for @refreshka/ispolzuite-red-i-pokupat-energiyu-mozhno-budet-na-30-deshevle
2018/03/05 12:55:20 @sonik voted for @process/zhurnal-psikhologiya-net-resursa-dvigatsya-dalshe-ili-kuda-propalo-vdokhnovenie
2018/03/05 12:55:20 @madrid voted for @just-life/zhizn-pod-otkrytym-nebom-chast-2
2018/03/05 12:55:20 @dobryj.kit voted for @vasyl73/na-ulice-vesna-sneg-i-22-gradusa
2018/03/05 12:55:20 @ruta voted for @vp-magic-india/morskaya-arkheologiya-indii
2018/03/05 12:55:20 @gans91 voted for @chugoi/inogda-oni-vozvrashayutsya-rynok-vospominanii
2018/03/05 12:55:20 @vetal voted for @yura81/zarabotaite-v-steemit-com-ne-menee-20-rublei-za-kommentarii
2018/03/05 12:55:20 @vetal voted for @fyyf/75yfpx-krasnyi-ponedelnik
2018/03/05 12:55:21 @alchemy.inv voted for @alchemy.inv/yaponskie-svechi
...
```
