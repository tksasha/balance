# frozen_string_literal: true

RSpec.describe 'Backoffice::Cashes', type: :request do
  %w[uah usd eur].each do |currency|
    describe 'GET index.js' do
      before do
        create_list(:cash, 2, currency:)

        get "/#{ currency }/backoffice/cashes", xhr: true
      end

      it_behaves_like 'index.js'
    end

    describe 'GET new.js' do
      before { get "/#{ currency }/backoffice/cashes/new", xhr: true }

      it_behaves_like 'new.js'
    end

    describe 'POST create.js' do
      before { post "/#{ currency }/backoffice/cashes", params:, xhr: true }

      context 'with valid params' do
        let(:params) do
          {
            cash: {
              name: 'Bank',
              formula: '4 + 5',
              supercategory: 'bonds',
              favorite: true
            }
          }
        end

        let(:cash) { Cash.last }

        it_behaves_like 'create.js'

        it { expect(cash.name).to eq 'Bank' }

        it { expect(cash.sum).to eq 9 }

        it { expect(cash.supercategory).to eq 'bonds' }

        it { expect(cash.favorite).to be_truthy }

        it { expect(cash.currency).to eq currency }
      end

      context 'with invalid params' do
        let(:params) { { cash: { name: '' } } }

        it_behaves_like 'new.js'
      end
    end

    describe 'GET edit.js' do
      let(:cash) { create :cash }

      before { get "/#{ currency }/backoffice/cashes/#{ cash.id }/edit", xhr: true }

      it_behaves_like 'edit.js'
    end

    describe 'PATCH update.js' do
      let(:cash) { create(:cash, currency:) }

      before { patch "/#{ currency }/backoffice/cashes/#{ cash.id }", params:, xhr: true }

      context 'with valid params' do
        before { cash.reload }

        let(:name) { Faker::Lorem.word }

        let(:params) do
          {
            cash: {
              name:,
              supercategory: 'bonds',
              favorite: true
            }
          }
        end

        it_behaves_like 'update.js'

        it { expect(cash.name).to eq name }

        it { expect(cash.supercategory).to eq 'bonds' }

        it { expect(cash.favorite).to be_truthy }

        it { expect(cash.currency).to eq currency }
      end

      context 'with invalid params' do
        let(:params) { { cash: { name: '' } } }

        it_behaves_like 'edit.js'
      end
    end

    describe 'DELETE destroy.js' do
      let(:cash) { create :cash }

      before { delete "/#{ currency }/backoffice/cashes/#{ cash.id }", xhr: true }

      it_behaves_like 'destroy.js'

      it { expect { cash.reload }.to raise_error(ActiveRecord::RecordNotFound) }
    end
  end
end
