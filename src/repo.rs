use std::process::{Command, Stdio};

use anyhow::{ensure, Context, Result};

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

impl Repository {
    pub fn git_clone(&self) -> Result<()> {
        let status = Command::new("git")
            .arg("clone")
            .arg(format!("https://github.com/{}/{}", self.owner, self.name))
            .status()
            .context("Failed to setup process to clone repo with git")?;
        ensure!(status.success());
        Ok(())
    }
}

#[cfg(test)]
mod test {
    use std::fs;

    use anyhow::Result;

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

    #[test]
    fn git_clone() -> Result<()> {
        assert!(Repository {
            owner: String::from("gleich"),
            name: String::from("fgh")
        }
        .git_clone()
        .is_ok());
        fs::remove_dir_all("fgh").expect("Failed to remove cloned dir: fgh");
        assert!(!Repository {
            owner: String::from("gleich"),
            name: String::from("")
        }
        .git_clone()
        .is_ok());
        Ok(())
    }
}
