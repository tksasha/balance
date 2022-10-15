# frozen_string_literal: true

RSpec.describe 'Frontend/Cashes', type: :request do
  describe 'GET /frontend/cashes.js' do
    %w[uah usd eur].map do |currency|
      context "when currency is `#{ currency }`" do
        before { get "/#{ currency }/frontend/cashes", xhr: true }

        it_behaves_like 'index.js'
      end
    end
  end
end
