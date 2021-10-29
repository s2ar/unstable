[![Build Status](https://scrutinizer-ci.com/g/s2ar/unstable/badges/build.png?b=master)](https://scrutinizer-ci.com/g/s2ar/unstable/build-status/master)
![Go Report](https://goreportcard.com/badge/github.com/s2ar/unstable)
![Repository Top Language](https://img.shields.io/github/languages/top/s2ar/unstable)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/s2ar/unstable/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/s2ar/unstable/?branch=master)
![Lines of code](https://img.shields.io/tokei/lines/github/s2ar/unstable)
![Github Open Issues](https://img.shields.io/github/issues/s2ar/unstable)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/s2ar/unstable)
![Github Repository Size](https://img.shields.io/github/repo-size/s2ar/unstable)
![GitHub last commit](https://img.shields.io/github/last-commit/s2ar/unstable)
# Unstable service
## Описание
При запуске сохраняются данные с внешнего источника. 
Сервис имеет эндпоинт http://localhost:8083/api/team/top по которому в случайном 
порядке может отдать статус 200, 408 или 500
## Howto:
`make run`

Вывод в консоль:
![2021-10-29-2e6](https://user-images.githubusercontent.com/2817417/139493611-3333c98b-15d8-4dc3-b68f-786096d9406a.png)

