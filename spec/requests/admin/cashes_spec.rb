# frozen_string_literal: true

RSpec.describe 'Admin/Cashes' do
  describe 'POST create' do
    before { post '/admin/cashes', params: }

    context 'when params are valid' do
      let(:params) do
        {
          cash: {
            name: 'Cash #1',
            formula: '4.2 + 6.9',
            currency: 'eur',
            favorite: true
          }
        }
      end

      let(:cash) { Cash.last }

      it { is_expected.to redirect_to admin_cash_path(cash) }

      it { expect(cash.name).to eq 'Cash #1' }
      it { expect(cash.sum).to eq 11.1 }
      it { expect(cash.currency).to eq 'eur' }
      it { expect(cash).to be_a_favorite }
    end

    context 'when params are not valid' do
      let(:params) { { cash: { name: '' } } }

      it { is_expected.to render_template :new }
    end
  end
end
