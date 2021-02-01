# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  it_behaves_like 'index', '/cashes.js'

  it_behaves_like 'edit', '/cashes/54/edit.js' do
    before { create :cash, id: 54 }
  end

  it_behaves_like 'update', '/cashes/49.js' do
    let :valid_params do
      { cash: { name: 'Bank' } }
    end

    let :invalid_params do
      { cash: { name: '' } }
    end

    before { create :cash, id: 49 }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end
end
