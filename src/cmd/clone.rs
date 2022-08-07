use clap::ArgMatches;

use crate::repo::Repository;

pub fn run(args: &ArgMatches) {
    let repo = Repository::from(
        args.get_one::<String>("repository")
            .expect("Failed to get argument for repository")
            .as_str(),
    );
}
