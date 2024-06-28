{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = inputs:
    with inputs;
      flake-utils.lib.eachDefaultSystem (
        system: let
          pkgs = import nixpkgs {
            inherit system;

            config.allowUnfree = true;
          };
        in {
          devShells = with pkgs; {
            default = mkShell {
              packages = [
                go

                goose
                go-jet

                just

                postman
              ];
            };
          };
        }
      );
}
