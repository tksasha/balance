require 'rails_helper'

RSpec.describe CashesController, type: :controller do
  it { should be_a ActsAsRESTController }

  describe '#collection' do
    before { expect(Cash).to receive(:order).with(:name).and_return(:collection) }

    its(:collection) { should eq :collection }
  end

  describe '#resource_params' do
    before { expect(subject).to receive(:params).and_return(acp cash: { formula: '', name: '' }) }

    its(:resource_params) { should eq permit! formula: '', name: '' }
  end

  it_behaves_like :update do
    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end

  it_behaves_like :create do
    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end

  describe '#set_variant' do
    context do
      before { expect(subject).to receive(:params).and_return({}) }

      before { expect(subject).to_not receive(:request) }

      it { expect { subject.send :set_variant }.to_not raise_error }
    end

    context do
      before { expect(subject).to receive(:params).and_return({ report: '' }) }

      before { expect(subject).to_not receive(:request) }

      it { expect { subject.send :set_variant }.to_not raise_error }
    end

    context do
      before { expect(subject).to receive(:params).and_return({ report: '1' }) }

      before do
        #
        # subject.request.variant = :report
        #
        expect(subject).to receive(:request) do
          double.tap { |a| expect(a).to receive(:variant=).with(:report) }
        end
      end

      it { expect { subject.send :set_variant }.to_not raise_error }
    end
  end
end
