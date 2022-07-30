use clap::{Arg, Command};

pub fn setup() -> Command<'static> {
    Command::new("fgh")
        .version("1.0.0")
        .author("Matt Gleich <email@mattglei.ch>")
        .about("Automate the lifecycle and organization of your cloned GitHub repositories")
        .arg_required_else_help(true)
        .subcommand(
            Command::new("clone")
                .about("Clone a repository")
                .arg(Arg::new("repository").takes_value(true).index(1)),
        )
}
