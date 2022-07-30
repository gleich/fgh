#[derive(Debug, Default, PartialEq)]
pub struct Repository {
    pub owner: String,
    pub name: String,
}

impl From<&str> for Repository {
    /// Parse owner and name of repo from either format:
    /// 1. https://github.com/gleich/fgh
    /// 2. gleich/fgh
    fn from(s: &str) -> Self {
        let mut pair = s;
        let (owner, name);
        if s.starts_with("https://github.com/") {
            pair = pair.strip_prefix("https://github.com/").unwrap();
        }
        if pair.contains("/") {
            (owner, name) = pair.split_once("/").unwrap()
        } else {
            (owner, name) = ("", pair);
        }
        Repository {
            owner: owner.to_string(),
            name: name.to_string(),
        }
    }
}

#[cfg(test)]
mod test {
    use super::Repository;

    #[test]
    fn from_str() {
        let repo = Repository {
            owner: String::from("gleich"),
            name: String::from("fgh"),
        };
        assert_eq!(Repository::from("https://github.com/gleich/fgh"), repo);
        assert_eq!(Repository::from("gleich/fgh"), repo);
        // Shouldn't be equal for now as owner is not detected
        assert_ne!(Repository::from("fgh"), repo);
    }
}
