use serde::{Deserialize, Serialize};

#[derive(Serialize)]
pub struct WorldRequest {
    pub width: u32,
    pub height: u32,
    pub seed: u32
}

#[derive(Deserialize)]
pub struct WorldResponse {
    pub width: u32,
    pub height: u32,
    pub grid: Vec<Vec<String>>,
}

pub async fn generate_world() -> Result<WorldResponse, reqwest::Error> {
    let client = reqwest::Client::new();

    let response = client
        .post("http://localhost:3000/api/v1/worlds")
        .json(&WorldRequest {
            width: 20,
            height: 10,
            seed: 123
        })
        .send()
        .await?
        .error_for_status()?
        .json::<WorldResponse>()
        .await?;

    Ok(response)
}