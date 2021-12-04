# frozen_string_literal: true

RSpec.describe 'Backoffice::Cashes', type: :request do
  it_behaves_like 'new', '/backoffice/cashes/new.js'

  it_behaves_like 'create', '/backoffice/cashes.js' do
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

  it_behaves_like 'update', '/backoffice/cashes/49.js' do
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

  it_behaves_like 'destroy', '/backoffice/cashes/49.js' do
    before { create :cash, id: 49 }

    let(:success) { -> { should render_template :destroy } }
  end

  describe 'GET /index.js' do
    before { create_list :cash, 2 }

    before { get '/backoffice/cashes', xhr: true }

    it_behaves_like 'index.js'
  end
end
