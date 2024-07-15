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

    describe 'POST create.js' do
      context 'when params are invalid' do
        let(:params) { { category: { name: '' } } }

        before { post "/#{ currency }/categories", params:, xhr: true }

        it_behaves_like 'new.js'
      end

      context 'when params are valid' do
        let(:name) { Faker::Commerce.product_name }

        let(:params) { { category: { name: } } }

        let(:category) { Category.last }

        before { post "/#{ currency }/categories", params:, xhr: true }

        it_behaves_like 'create.js'

        it { expect(category.name).to eq name }

        it { expect(category.currency).to eq currency }
      end
    end
  end
end
