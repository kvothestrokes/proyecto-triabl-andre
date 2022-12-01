# API Buscadora de canciones
Esta es una API Restful que realiza consolida búsquedas en las siguientes apis de canciones: https://itunes.apple.com, http://api.chartlyrics.com
## Instalación
```
git clone 
cd 
docker compose up -d
```
## Uso
###  /get-token - GET
Obtiene el token para el uso de la api
Ejemplo respuesta:
```
{
	"status":true,
	"Error":"",
	"data"{
		"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlLCJleHBpcmUiOjE2Njk5NDkxNDB9.b-E3NOZE593vCgJGUzcdjyB0JnuCHcxtvfWyJoIT9MA"
	}
}
```
Tipo de datos de respuesta:
| Parámetro | Tipo | Descripción|
| --- | ---| ---| 
| status | bool | estatus de la petición
| Error | string | error generado (si existe)
| token | string | token nuevo creado


### /search  - GET
Parámetros de entrada:
| Parámetro | Tipo | Descripción |
| --- | ---| --- |
| name | string | Nombre de la canción
| token | string | Token generado en ruta /get-token
| artist | string | Nombre del artista
| album | string | Nombre del álbum

ejemplo de ruta:
```
http://localhost/search?name=reptilia&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlLCJleHBpcmUiOjE2Njk5MTgzODh9.6LUpMnzQGPWgZLWU9fzsjzKN78Vm0mVYmpWnmwOS5rE&artist=strokes
```
ejemplo de respuesta:
```
{
	"status":true,
	"Error":"",
	"data":[
		{
			"id":"ireptilia296261076",
			"name":"reptilia",
			"artist":"the strokes",
			"duration":"03:36",
			"album":"",
			"artwork":"https://is4-ssl.mzstatic.com/image/thumb/Video113/v4/50/54/33/50543349-e382-c3d9-3645-85d5e36a1179/dj.vbvpugtf.jpg/100x100bb.jpg",
			"price":"USD 1.99",
			"origin":"itunes"
			},
		{
			"id":"mreptilia15334",
			"name":"reptilia",
			"artist":"the strokes",
			"duration":"00:00",
			"album":"",
			"artwork":"",
			"price":" 0.00",
			"origin":"chartlyrics"}
		]
	}
```
Datos de respuesta:
| Parámetro | Tipo | Descripción |
| --- | ---| --- |
| id | string | Id único de canción|
| name | string | Nombre de la canción|
| artist | string | Nombre del artista
| album | string | Nombre del álbum|
| duration | string | Tiempo de la canción en minutos|
| artwork | string | url de la imágen del álbum |
|price | string | precio y moneda de la canción |
| origin | string | API de origen de los datos (itunes o chartlyrics) |







