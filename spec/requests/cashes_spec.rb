# frozen_string_literal: true

RSpec.describe 'Cashes', type: :request do
  it_behaves_like 'new', '/cashes/new.js'

  it_behaves_like 'destroy', '/cashes/49.js' do
    before { create :cash, id: 49 }

    let(:success) { -> { should render_template :destroy } }
  end

  it_behaves_like 'create', '/cashes.js' do
    let :valid_params do
      {
        cash: {
          name: 'Bank',
          formula: '4 + 5',
        }
      }
    end

    let :invalid_params do
      { cash: { name: '' } }
    end

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
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
