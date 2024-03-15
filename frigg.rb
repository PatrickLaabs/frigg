# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Frigg < Formula
  desc ""
  homepage "https://github.com/PatrickLaabs/frigg"
  version "1.1.1"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.1.1/frigg_1.1.1_darwin_arm64.tar.gz"
      sha256 "f33c85ece6e1265e17f31003f3c9506b5438067bc2d13dc360088afe7a54867c"

      def install
        bin.install "frigg"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.1.1/frigg_1.1.1_darwin_amd64.tar.gz"
      sha256 "b4e01545da0c55e7886635b0cc4000e78b61c0d515965969e1583a3881b352c0"

      def install
        bin.install "frigg"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.1.1/frigg_1.1.1_linux_arm64.tar.gz"
      sha256 "4dae358e77fbdb91725fb5683140d2a1358ee27011981de9f5646e2406596e61"

      def install
        bin.install "frigg"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/PatrickLaabs/frigg/releases/download/1.1.1/frigg_1.1.1_linux_amd64.tar.gz"
      sha256 "e25e16d2b7fe03af50f52b5a32f23322bde5db2aa4cf4d9020600b99f395825d"

      def install
        bin.install "frigg"
      end
    end
  end
end
