mod cli;
mod cmd;
mod repo;

fn main() {
    let matches = cli::setup().get_matches();
    match matches.subcommand() {
        Some(("clone", args)) => cmd::clone::run(args),
        _ => unreachable!(),
    }
}
