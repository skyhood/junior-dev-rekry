# Junior Software Developeriksi Duunitorille!

Moikka ja onnea tähän vaiheeseen pääsemisestä, olet jo osa pientä harvojen ja valittujen joukkoa! Alta löydät meidän Junior Software Developer -työnhakuputkeen liittyvän ennakkotehtävän. Aloitetaan muutamalla yleisen tason infolla.

1. Voit valita vahvuuksiesi pohjalta tehdä joko frontend tai backend -tehtävän, molempia ei siis ole tarkoitus tehdä!
2. Suunnittelimme tehtävän kestämään 2-5 h, emmemekä odota, että käytät tähän aikaa sen enempää.
3. Emme odota sinun olevan valmis ammattilainen, eikä koodisikaan ole täydellistä (eihän se koskaan ole). Kuitenkin, etsimme tyyppiä joka pystyy jo sopivasti tuettuna tekemään itsenäistä työtä, ja tätä pyrimme tehtävällä tunnistamaan. Arviointikriteereistä voit lukea lisää tehtäväkuvausten alta.


## Tehtävä – Frontend

Löydät `jobs.json` -tiedostosta löydät vinon pinon työpaikkailmotusdataa (1000 kpl jos ihan tarkkoja ollaan). Tee näistä työpaikkojen listaussivu.

Sivulla tulisi olla mahdollista nähdä olemassa olevat avoimet työpaikat. Sivulla voi olla jos minkä sorttisia ominaisuuksia filttereistä animointeihin tai sivutuksesta upeaan ulkoasuun, vain taivas (ja aika) on rajana.

Voit tehdä sivuston ilmeestä sellaisen kuin haluat. Inspiraatiota voit ottaa esimerkiksi meidän [hakutulossivuilta](https://duunitori.fi/tyopaikat). Laitoimme liitteeksi myös css-tiedoston, josta löytyy hieman tukea tähän :). Visuaalinen ilme sinäänsä ei ole pakollinen arviointikriteeri, mutta odotamme toki sivun palvelevan tarkoitustaan.


### Arviointi
Yleiset arviointikriteerit löydät tehtäväkuvausten alta.

Tässä tehtävässä voit valita, haluatko keskittyä työssäsi _visualisuuteen_, _tehokkuuteen_, _saavutettavuuteen_ ja/vai _käytettävyyteen_. Kerrothan meille, mihin 1-2 asiaan näistä olet keskittynyt, jotta osaamme arvioida tuotostasi oikeasta perspektiivistä :)


#### Visuaalisuus
Meillä devaajat ja designerit etsivät parhaita tapoja ratkaista kysymyksiä yhdessä. Kun devi ja design löytävät yhteisen sävelen ja tuloksena on jotakin upeaa.

#### Tehokkuus
Meillä, kuten tässä tehtävässä, dataa on valtavasti. Kaikki eivät kuitenkaan etsi töitä uusimmilla vempeleillä ja nopeimmilla nettiyhteyksillä, joten tehokkuuden huomioimainen sivustomme kehityksessä.

#### Saavutettavuus
Me haluamme tehdä työelämästä monimuotoista, joten koodin saavutettavuus ja esteellisyyden poistaminen on yksi tavoitteistamme.


#### Käytettävyys
Ärsyttävät palvelut ovat ärsyttäviä. Emme halua lisätä ärsytystä, vaan tehdä palvelukokemustamme käytännöllisen, toimivan ja kenties jopa ilahduttavan.



## Tehtävä – Backend

Löydät `jobs.json` -tiedostosta löydät vinon pinon työpaikkailmotusdataa (1000 kpl jos ihan tarkkoja ollaan). Tee näistä API. Rajapintaa käyttää nettisivu, joka listaa ja filtteröi avoimia työpaikkailmoituksia, sekä näyttää ilmoitusten tietoja.

Voit hyödyntää vapaasti valitsemaasi koodikieltä ja apukirjastoja. Jos kaipaat instpiraatiota stäkin valintaan, meillä käytetään sekä Pythonia Djangon kera että Gota.


### Ominaisuuksia

Tässä muutama esimerkkiominaisuus, joita voit implementoida. Kaikkea ei tarvitse tehdä, vaan valitse mitä haluat ja ehdit annetussa ajassa tekemään.

#### Haku
Primitiivisen hakutoiminnon implementointi. Haku arvatenkin on meillä aivan palvelun ytimessä, sillä ihmisiä kiinnostaa luonnollisesti omat (ja naapurin) työt.

#### Tilastodata
Datan aggregointi kiinnostaviksi luvuiksi tai dataseteiksi. Mitä, missä, milloin, miksi, miten paljon? Esimerkkidatan pohjalta mahdollisuudet on toki hieman rajalliset, mutta meillä monenlaista dataa on valtavasti ja tätä koitamme hyödyntää monin tavoin.

#### Sivutus
Implementoi rajapinnan palauttamiin tuloksiin jonkunlainen sivutus. No johan on työpaikkoja, meillä niitä löytyy joka hetki n. 50 000. Ihan jokaista ei kannata kerralla silmille läväyttää.

#### Top-listat
Implementoi rajapinta erilaisille kiinnostaville top-listoille. ”7 kuuminta työnantajaa – tutustu näihin jos haluat työllistyä?”, ”Muutto edessä? katso 12 kuntaa, joista löytyy töitä!”. Meillä kahlaillaan jatkuvasti dataa, jotta voimme tuottaa kiinnostavia sisältöjä.

#### Yksittäinen ilmoitus
Implementoi rajapinta, joka palauttaa yksittäisen ilmoituksen sisällön. Vaikka listat on kivoja, vielä kivempaa on kun sieltä löytää jotain itseään kiinnostavaa.


### Arviointi
Yleiset arviointikriteerit löydät tehtäväkuvausten alta.

Tässä tehtävässä kiinnitä erityistä huomiota koodin tehokkuuteen. Huomioithan myös käyttötarkoituksen – tarvitseeko nettisaitti kerralla 1000 työpaikkaa kaikkine tietoineen? Meillä, kuten tässä tehtävässä, dataa on valtavasti. Kaikki eivät kuitenkaan etsi töitä uusimmilla vempeleillä ja nopeimmilla nettiyhteyksillä, ja tämän pyrimme huomioimaan sivustomme kehityksessä.


## Arviointikriteerit

Pyrimme selkeyteen ja ymmärrettävyyteen arvioinnissamme. Tehtäväkohtaisten kuvausten lisäksi kiinnitämme erityistä huomiota seuraaviin asioihin

### Koodin toimivuus
Homman tulee lähtökohtaisesti toimia. Fronttituotosta tutkimme lähtökohtaisesti uusimmalla Chromella. Muiden selaimien tukemisesta ei tarvitse huolehtia. Liitäthän mukaan selkeät ohjeet, joiden perusteella saamme softasi pyörimään, meillä ei valitettavasti ole aikaa puutteellisesti toimivan koodin debuggaamiseen, käytämme sen mielummin oman koodimme bugien liiskaamiseen :)

### Koodin ymmärrettävyys
Meillä työskennellään porukalla samojen asioiden parissa, joten muidenkin on ymmärrettävä kirjoitettua koodia. Noudata hyviä käytäntöjä nimeämisten, kommenttien, dokumentoinnin ym suhteen. Olisi valitettavaa, jos koodisi toimisi upeasti mutta emme pystyisi ymmärtämään, miksi.

### Koodin uudelleenkäytettävyys
Jotta meidän ei tarvitse kirjoittaa samanlaisia koodirivejä uudestaan ja uudestaan, koitamme tunnistaa uudelleenkäytettävät osat koodistamme. Tällaista perspektiiviä omaa tekemistä kohtaan arvostamme myös sinulta. [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) on mukava periaate (Joskin ei sitäkään kannata ylisuorittaa, been there.)

### Versionhallinta
Meillä käytetään gittiä, ja osaamisen vähintään gitin perusteista katsotaan eduksi. Lopputuloksen toimittaminen ymmärrettävän git-historian kera on mahdollinen lisäpisteen paikka, joskin ei vaatimus.

## jobs.json

JSON-tiedosto näyttää tältä

```
{
  "jobs": [
    {
      "heading": string
      "date_posted": string
      "slug": string
      "municipality_name": string
      "export_image_url": string
      "company_name": string
      "descr": string
      "latitude": int
      "longitude": int
      "salary": {
        "type": int
        "min": int
        "max": int
      }
    }, ...
  ]
}
```


## Valmiin tuotoksen toimittaminen meille päin

Tadaa, valmista! Voit palauttaa tehtävän haluamallasi tavalla. Itse käytämme GitHubia, ja uusi yksityinen repository on oikein toimiva tapa toimittaa valmis lopputulos. Huomioithan kuitenkin muutaman asian

1. *Älä forkkaa tätä repositoryä tai tee pull requestia suoraan tänne!* Muutoin työsi näkyy julkisesti kaikille, ja olisi ikävää jos kanssahakijasi kopioisi työtäsi.
2. *Jos hyödynnät GitHubia tai vastaavaa, älä tee repositorystäsi julkista* samasta syystä kuin yllä
3. Toimitathan työsi mukana mahdollisimman selkeät ohjeet miten pystymme tuotostasi kokeilemaan. Meillä ei valitettavasti ole aikaa ratkoa bugeja koodistasi :)
