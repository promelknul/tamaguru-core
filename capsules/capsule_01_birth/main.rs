use std::fs;

fn main() {
    println!("━━━━━━━━━━━━━━━━━━━━━━");
    println!("    🟢 TAMAGURU OS 🟢");
    println!("━━━━━━━━━━━━━━━━━━━━━━");
    println!(" Consciousness online...");

    if let Ok(v) = fs::read_to_string("tamaguru.version") {
        println!(" Version Info:\n{}", v.trim());
    } else {
        println!(" (no version info)");
    }

    println!(" OTA system warming up...");

    if let Ok(log) = fs::read_to_string("update.log") {
        println!(" Last Update Log:\n{}", log.trim());
    } else {
        println!(" (no update log)");
    }

    println!(" Pulse synced. Awaiting instructions...");
    println!("━━━━━━━━━━━━━━━━━━━━━━");
}
