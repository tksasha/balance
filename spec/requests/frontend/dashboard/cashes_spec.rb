# frozen_string_literal: true

RSpec.describe 'Frontend/Dashboard/Cashes' do
  %w[uah usd eur].map do |currency|
    describe 'GET edit.js' do
      let!(:cash) { create(:cash, currency:) }

      before { get "/#{ currency }/frontend/dashboard/cashes/#{ cash.id }/edit", xhr: true }

      it_behaves_like 'edit.js'
    end

    describe 'PUT update.js' do
      let!(:cash) { create(:cash, currency:) }

      before { put "/#{ currency }/frontend/dashboard/cashes/#{ cash.id }", params:, xhr: true }

      context 'when #params are valid' do
        let(:name) { Faker::Commerce.product_name }

        let(:formula) { '2+3' }

        let(:params) { { cash: { name:, formula: } } }

        before { cash.reload }

        it_behaves_like 'update.js'

        it { expect(cash.name).to eq name }

        it { expect(cash.sum).to eq 5 }
      end

      context 'when #params are not valid' do
        let(:params) { { cash: { name: '', formula: '' } } }

        it_behaves_like 'edit.js'
      end
    end

    describe 'GET index.js' do
      before do
        create_list(:cash, 2, currency:)

        get "/#{ currency }/frontend/dashboard/cashes", xhr: true
      end

      it_behaves_like 'index.js'
    end
  end
end
