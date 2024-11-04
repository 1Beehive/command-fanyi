class Fanyi < Formula
  desc "A command-line tool for translation"
  homepage "https://github.com/jeffcail/command-fanyi" # 替换为你的项目主页
  url "https://github.com/jeffcail/command-fanyi", :using => :nounzip # 替换为你的 fanyi 程序的路径
  version "0.0.1" # 替换为你的版本号

  def install
    bin.install "fanyi"
  end

  test do
    system "#{bin}/fanyi", "--version" # 这里可以替换为你希望测试的命令
  end
end