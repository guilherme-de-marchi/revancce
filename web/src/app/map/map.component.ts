import { Component, OnInit, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

import Map from 'ol/Map';
import View from 'ol/View';
import Feature from 'ol/Feature';
import Overlay from 'ol/Overlay.js';
import Point from 'ol/geom/Point.js';
import {Icon, Style} from 'ol/style.js';
import {OGCMapTile, Vector as VectorSource} from 'ol/source.js';
import {Tile as TileLayer, Vector as VectorLayer} from 'ol/layer.js';
import * as style from 'ol/style';
import * as proj from 'ol/proj';
import * as geom from 'ol/geom';
import * as interaction from 'ol/interaction';
import * as layer from 'ol/layer';
import * as source from 'ol/source';
import OSM from 'ol/source/OSM';
import * as olProj from 'ol/proj';

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss']
})
export class MapComponent implements OnInit {
  @Input() coord!: [number, number]
  public map!: Map

  ngOnInit(){
    this.map = new Map({
      interactions: interaction.defaults({mouseWheelZoom:false}),
      target: 'map',
      layers: [
        new layer.Tile({
          source: new source.OSM(),
        }),
        new layer.Vector({
          source: new source.Vector({
            features: [
              new Feature({
                geometry: new geom.Point(proj.fromLonLat(this.coord)),
                name: 'Somewhere near Nottingham',
              })
            ]
          }),
          style: new style.Style({
            image: new style.Icon({
              anchor: [0.5, 46],
              anchorXUnits: 'fraction',
              anchorYUnits: 'pixels',
              src: 'https://ucarecdn.com/4b516de9-d43d-4b75-9f0f-ab0916bd85eb/marker.png'
            })
          })
        })
      ],
      view: new View({
        center: proj.fromLonLat(this.coord),
        zoom: 17
     })    
    });
  }
}
