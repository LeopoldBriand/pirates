use bevy::{DefaultPlugins, app::{App, Startup, Update}, camera::Camera2d};
use bevy::ecs::system::{Commands};
use crate::{internals::tilemap::TileMap, rendering::tilemap::{load_tile_textures, spawn_tilemap}};

mod http;
mod internals;
mod rendering;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, load_tile_textures)
        .add_systems(Startup, setup)
        .add_systems(Update, spawn_tilemap)
        .run();
}

fn setup(mut commands: Commands) {
    commands.spawn(Camera2d);

    let runtime = tokio::runtime::Runtime::new().unwrap();

    let world = runtime
        .block_on(http::world::generate_world())
        .expect("Failed to fetch map");
    
    let tilemap = TileMap::from(world);
    commands.insert_resource(tilemap);
}