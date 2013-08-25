require "fakeredis/rspec"
require "lita/rspec"

require "lita/handlers/google_images"

describe Lita::Handlers::GoogleImages, lita_handler: true do
  it { routes_command("/image me foo").to(:fetch) }
  it { routes_command("/image foo").to(:fetch) }
  it { routes_command("/img foo").to(:fetch) }
  it { routes_command("/img me foo").to(:fetch) }

  describe "fetch" do
    let(:response) { double("Faraday::Response") }

    before do
       allow_any_instance_of(Faraday::Connection).
         to receive(:get).and_return(response)
    end

    it "replies with a matching image URL on success" do
      url = "http://example.com/path/to/image.jpg"

      allow(response).to receive(:body).and_return('{
          "responseStatus": 200,
          "responseData": {
            "results": [
              { "unescapedUrl": "' + url + '" }
            ]}}')

      send_message("/image carl")

      expect(replies.last).to eq(url + '#.png')
    end
  end
end
