use bevy::ecs::system::{Commands, Res};

use crate::internals::tilemap::TileMap;

use std::{collections::HashMap, fs, path::Path};
use bevy::prelude::*;

#[derive(Resource)]
pub struct TileTextures {
    pub textures: HashMap<String, Handle<Image>>,
}

pub fn load_tile_textures(
    asset_server: Res<AssetServer>,
    mut commands: Commands,
) {
    let tiles_dir = Path::new("assets/tiles");

    let mut textures = HashMap::new();

    let entries = fs::read_dir(tiles_dir)
        .expect("Failed to read assets/tiles");

    for entry in entries {
        let path = entry
            .expect("Failed to read directory entry")
            .path();

        if path.extension().and_then(|ext| ext.to_str()) != Some("png") {
            continue;
        }

        let Some(filename) = path.file_stem().and_then(|name| name.to_str()) else {
            continue;
        };

        let asset_path = format!("tiles/{}.png", filename);

        let handle = asset_server.load(asset_path);

        textures.insert(filename.to_string(), handle);
    }

    commands.insert_resource(TileTextures { textures });
}

pub fn spawn_tilemap(
    mut commands: Commands,
    tilemap: Res<TileMap>,
    textures: Res<TileTextures>,
) {
    const TILE_SIZE: f32 = 32.0;

    for y in 0..tilemap.height {
        for x in 0..tilemap.width {
            let tile_id = tilemap.get(x, y);

            let Some(texture) = textures.textures.get(&tile_id) else {
                warn!("Unknown tile: {tile_id}");
                continue;
            };

            commands.spawn((
                Sprite {
                    image: texture.clone(),
                    ..default()
                },
                Transform::from_xyz(
                    x as f32 * TILE_SIZE,
                    -(y as f32) * TILE_SIZE,
                    0.0,
                ),
            ));
        }
    }
}