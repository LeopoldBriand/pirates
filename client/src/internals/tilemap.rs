// src/world/tilemap.rs

use bevy::ecs::resource::Resource;

#[derive(Debug, Resource)]
pub struct TileMap {
    pub width: usize,
    pub height: usize,
    pub tiles: Vec<String>,
}

impl TileMap {
    pub fn new(width: usize, height: usize, tiles: Vec<String>) -> Self {
        assert_eq!(tiles.len(), width * height);

        Self {
            width,
            height,
            tiles,
        }
    }

    pub fn get(&self, x: usize, y: usize) -> String {
        self.tiles[y * self.width + x].clone()
    }
}

impl From<crate::http::world::WorldResponse> for TileMap {
    fn from(map: crate::http::world::WorldResponse) -> Self {
        let tiles = map.grid.into_iter().flatten().collect();

        Self::new(
            map.width as usize,
            map.height as usize,
            tiles,
        )
    }
}