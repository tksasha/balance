# frozen_string_literal: true

RSpec.describe 'Cashes' do
  Item.currencies.keys.map do |currency|
    describe 'GET /cashes.js' do
      before { get "/#{ currency }/cashes", xhr: true }

      it_behaves_like 'index.js'
    end

    describe 'GET edit.js' do
      let!(:cash) { create(:cash, currency:) }

      before { get "/#{ currency }/cashes/#{ cash.id }/edit", xhr: true }

      it_behaves_like 'edit.js'
    end

    describe 'PUT /frontend/cashes/:id.js' do
      let!(:cash) { create(:cash, currency:) }

      before { put "/#{ currency }/cashes/#{ cash.id }", params:, xhr: true }

      context 'when params are valid' do
        let(:name) { Faker::Commerce.product_name }

        let(:params) do
          { cash: { name:, formula: '13.13 + 17.10' } }
        end

        before { cash.reload }

        it_behaves_like 'update.js'

        it { expect(cash.name).to eq name }
        it { expect(cash.sum).to eq 30.23 }
      end

      context 'when params are invalid' do
        let(:params) { { cash: { name: '' } } }

        it_behaves_like 'edit.js'
      end
    end
  end
end
