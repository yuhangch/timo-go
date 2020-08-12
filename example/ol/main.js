import 'ol/ol.css';
import Map from 'ol/Map';
import OSM from 'ol/source/OSM';
import MVT from 'ol/format/MVT';
import TileJSON from 'ol/source/TileJSON';
import TileLayer from 'ol/layer/Tile';
import VectorTileLayer from 'ol/layer/VectorTile';
import VectorTileSource from 'ol/source/VectorTile';
import View from 'ol/View';

var map = new Map({
  layers: [new TileLayer({
    source: new OSM(),
  }),
    new VectorTileLayer({
      declutter: true,
      source: new VectorTileSource({
        attributions:
          '© <a href="https://timo.yhch.cloud">' +
          'TIMO</a>',
        format: new MVT(),
        url:
          'https://timo.yhch.cloud/tiles/countries/' +
          '{z}/{x}/{y}.pbf?columns=fid,NAME'
      })
    }) ],
  target: 'map',
  view: new View({
    center: [0, 0],
    zoom: 2,
  }),
});
map.on('pointermove', showInfo);

var info = document.getElementById('info');
function showInfo(event) {
  var features = map.getFeaturesAtPixel(event.pixel);
  if (features.length == 0) {
    info.innerText = '';
    info.style.opacity = 0;
    return;
  }
  var properties = features[0].getProperties();
  info.innerText = JSON.stringify(properties, null, 2);
  info.style.opacity = 1;
}