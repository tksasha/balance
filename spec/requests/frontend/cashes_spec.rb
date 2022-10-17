# frozen_string_literal: true

# rubocop:disable RSpec/NestedGroups
RSpec.describe 'Frontend/Cashes', type: :request do
  %w[uah usd eur].map do |currency|
    context "when currency is `#{ currency }`" do
      describe 'GET /frontend/cashes.js' do
        before { get "/#{ currency }/frontend/cashes", xhr: true }

        it_behaves_like 'index.js'
      end

      describe 'GET /frontend/cashes/:id/edit.js' do
        let!(:cash) { create(:cash, currency:) }

        before { get "/#{ currency }/frontend/cashes/#{ cash.id }/edit", xhr: true }

        it_behaves_like 'edit.js'
      end

      describe 'PUT /frontend/cashes/:id.js' do
        let!(:cash) { create(:cash, currency:) }

        before { put "/#{ currency }/frontend/cashes/#{ cash.id }", params:, xhr: true }

        context 'when params are valid' do
          let(:params) { { cash: { name: Faker::Commerce.product_name, formula: '13.13 + 17.10' } } }

          before { cash.reload }

          it_behaves_like 'update.js'

          it { expect(cash.sum).to eq 30.23 }
        end

        context 'when params are invalid' do
          let(:params) { { cash: { name: '' } } }

          it_behaves_like 'edit.js'
        end
      end
    end
  end
end
# rubocop:enable RSpec/NestedGroups
