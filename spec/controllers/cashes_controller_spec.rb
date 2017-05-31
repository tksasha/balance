require 'rails_helper'

RSpec.describe CashesController, type: :controller do
  describe 'create.js' do
    let(:cash) { double }

    let(:params) { { cash: { name: 'Food', formula: '43.28 + 18.02' } } }

    before { expect(Cash).to receive(:new).with(permit! params[:cash]).and_return(cash) }

    before { expect(cash).to receive(:save!) }

    before { post :create, params: params, format: :js }

    it { should render_template :create }
  end

  describe 'edit.js' do
    before { get :edit, xhr: true, params: { id: 47 }, format: :js }

    it { should render_template :edit }
  end

  describe '#resource' do
    before { expect(subject).to receive(:params).and_return({ id: 31 }) }

    before { expect(Cash).to receive(:find).with(31).and_return(:resource) }

    its(:resource) { should eq :resource }
  end

  describe 'update.js' do
    let(:cash) { double }

    let(:params) { { cash: { name: 'Food', formula: '43.28 + 18.03' }, id: 1 } }

    before { subject.instance_variable_set :@cash, cash }

    before { expect(cash).to receive(:update!).with(permit! params[:cash]) }

    before { patch :update, params: params, format: :js }

    it { should render_template :update }
  end

  describe 'destroy.js' do
    let(:cash) { double }

    before { subject.instance_variable_set :@cash, cash }

    before { expect(cash).to receive(:destroy).and_return(true) }

    before { delete :destroy, params: { id: 1 }, format: :js }

    it { should render_template :destroy }
  end

  it_behaves_like :index

  describe '#collection' do
    before { expect(Cash).to receive(:order).with(:name).and_return(:collection) }

    its(:collection) { should eq :collection }
  end

  it_behaves_like :new

  describe '#initialize_resource' do
    before { expect(Cash).to receive(:new).and_return(:resource) }

    its(:initialize_resource) { should eq :resource }
  end
end
