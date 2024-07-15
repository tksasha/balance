# frozen_string_literal: true

RSpec.describe 'Categories' do
  %w[uah usd eur].each do |currency|
    describe 'GET index.js' do
      before do
        create_list(:category, 2, currency:)

        get "/#{ currency }/categories", xhr: true
      end

      it_behaves_like 'index.js'
    end

    describe 'GET new.js' do
      before { get "/#{ currency }/categories/new", xhr: true }

      it_behaves_like 'new.js'
    end
  end
end
