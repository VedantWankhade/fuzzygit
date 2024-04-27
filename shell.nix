let
    nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-23.11";
    pkgs = import nixpkgs { config = {}; overlays = []; };
in

pkgs.mkShell {
    packages = with pkgs; [
        go
        gotools
        go-tools
        gnumake
    ];

    shellHook = ''
        echo "============= fuzzygit ==================="
        echo "Entered fuzzygit dev environment."
        echo "------------------------"
        echo "Dependencies"
        go version
        make --version
        echo "------------------------"
        echo "Git status"
        git status
        export PS1="(fuzzygit)$PS1"
        echo "Setting up testing git directory"
        wd=`pwd`
        rm -rf /tmp/testdir
        mkdir /tmp/testdir
        make build-test
        cd /tmp/testdir
        git init
        git add -A
        git commit -m"test init"
        git branch test
        git branch dev
        git branch release
        cd "$wd"
    '';
}