require 'rails_helper'

RSpec.describe CategoriesController, type: :controller do
  describe '#resource_params' do
    let(:params) do
      { category: { name: 'Drinks', visible: true } }
    end

    before { expect(subject).to receive(:params).and_return(acp params) }

    its(:resource_params) { should eq permit! params[:category] }
  end

  it_behaves_like :index

  it_behaves_like :new

  it_behaves_like :create do
    let(:resource) { stub_model Category }

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :errors } }
  end

  it_behaves_like :edit

  it_behaves_like :update do
    let(:resource) { stub_model Category }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :errors } }
  end

  it_behaves_like :destroy do
    let(:success) { -> { should render_template :destroy } }
  end
end
