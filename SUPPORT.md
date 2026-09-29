# Getting help

There are three kinds of report, and each has its own channel. Picking the right
one gets you an answer faster and keeps security reports out of public view.

| You have                           | Go to                                                                                                                                     | Visibility |
| ---------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| A question about using soroauth    | a [blank issue](https://github.com/soroauth/soroauth-go/issues/new) labelled `question`                                                   | public     |
| A bug that is not a security issue | the [bug report form](https://github.com/soroauth/soroauth-go/issues/new?template=bug.yml)                                                | public     |
| A security vulnerability           | [GitHub Security Advisories](https://github.com/soroauth/soroauth-go/security/advisories/new), as described in [SECURITY.md](SECURITY.md) | private    |

## Questions

Read the [README](README.md), [ARCHITECTURE.md](ARCHITECTURE.md) and the guides
in [docs/](docs/) first; most usage questions (which credential arm simulation
returned, why the second simulation pass is needed, how expiration is counted)
are answered there.

If they are not, open an issue and say what you tried, what you expected, and
what happened instead. Include the entry XDR (base64) when the question is about
a specific entry, and the output of `soroauth inspect --entry <b64>`. Never
paste a secret seed: nothing soroauth does needs one to diagnose, and a seed
posted in an issue must be treated as compromised.

The repository has no discussion forum or chat channel. Questions about the
Stellar protocol itself, rather than this library, are better asked in the
[Stellar developer community](https://developers.stellar.org/).

## Bugs

Use the bug report form. A failing test or a command with its real output is the
most useful thing you can attach. See [CONTRIBUTING.md](CONTRIBUTING.md) if you
want to send the fix as well.

## Security reports

**Do not open a public issue.** Anything that could make soroauth produce a
signature over the wrong payload, write it to the wrong credential node, accept
an entry it should refuse, or leak a secret is a security report. Follow
[SECURITY.md](SECURITY.md), which defines the scope and the private reporting
channel. If you are unsure whether something counts, report it privately.
