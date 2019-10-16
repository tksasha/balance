# frozen_string_literal: true

RSpec.describe 'Categories', type: :request do
  let(:resource) { stub_model Category }

  let(:params) do
    {
      category: {
        name: nil,
        income: nil,
        visible: nil
      }
    }
  end

  let(:resource_params) { acp(params).require(:category).permit! }

  it_behaves_like 'new', '/categories/new.js'

  it_behaves_like 'edit', '/categories/49/edit.js' do
    before { expect(Category).to receive(:find).with('49').and_return(resource) }
  end

  it_behaves_like 'create', '/categories.js' do
    before { expect(Category).to receive(:new).with(resource_params).and_return(resource) }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end

  it_behaves_like 'update', '/categories/25.js' do
    before { expect(Category).to receive(:find).with('25').and_return(resource) }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end
end
