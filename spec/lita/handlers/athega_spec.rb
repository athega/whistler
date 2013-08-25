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
