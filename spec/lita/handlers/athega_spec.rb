require "fakeredis/rspec"
require "lita/rspec"

require "lita/handlers/athega"

describe Lita::Handlers::Athega, lita_handler: true do
  it { routes_http(:get, "/").to(:whistler) }

  it { routes("/athegian foo").to(:athegian) }

  describe "athegian" do
    before do
      Lita::Handlers::Athega.any_instance.
        should_receive(:get_athegian_data).
        with("foo").
        and_return({
          'medium_image_url' => 'foo.jpg',
          'name'             => 'bar',
          'position'         => 'baz'
        })
    end

    it "returns athegian data" do
      send_message("/athegian foo")

      replies.should == [
        "foo.jpg",
        "bar",
        "baz"
      ]
    end
  end

  it { routes("/xkcd foo").to(:xkcd) }

  describe "xkcd" do
    it "returns the latest xkcd comic by default" do
      Lita::Handlers::Athega.any_instance.
        should_receive(:get_latest_xkcd_data).
        and_return({
          'img'   => 'foo',
          'title' => 'bar',
          'alt'   => 'baz'
        })

      send_message('/xkcd')

      replies.should == ['foo', 'bar', 'baz']
    end

    it "returns the specified xkcd comic" do
      Lita::Handlers::Athega.any_instance.
        should_receive(:get_xkcd_data).
        with("123").
        and_return({
          'img'   => '123_foo',
          'title' => '123_bar',
          'alt'   => '123_baz'
        })

      send_message('/xkcd 123')

      replies.should == ['123_foo', '123_bar', '123_baz']
    end
  end

  it { routes("/aww").to(:aww) }

  describe "aww" do
    before do
      Lita::Handlers::Athega.any_instance.
        should_receive(:get_random_reddit_field).
        with("aww/top.json?limit=10", :url).
        and_return('foo.jpg')
    end

    it "returns a cute image" do
      send_message("/aww")
      replies.should == ['foo.jpg']
    end
  end
end
