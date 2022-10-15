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

  describe 'GET /frontend/cashes/:id/edit.js' do
    %w[uah usd eur].map do |currency|
      context "when currency is `#{ currency }`" do
        let!(:cash) { create(:cash, currency:) }

        before { get "/#{ currency }/frontend/cashes/#{ cash.id }/edit", xhr: true }

        it_behaves_like 'edit.js'
      end
    end
  end
end
