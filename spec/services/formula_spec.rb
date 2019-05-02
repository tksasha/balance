require 'rails_helper'

RSpec.describe Formula do
  describe '.calculate' do
    let(:string) { nil }

    subject { Formula.calculate string }

    it { should eq 0.0 }

    it { should be_a BigDecimal }

    context do
      let(:string) { '2+2' }

      it { should eq 4.0 }
    end

    context do
      let(:string) { '10-8' }

      it { should eq 2.0 }
    end

    context do
      let(:string) { '3*4' }

      it { should eq 12.0 }
    end

    context do
      let(:string) { 'string(2)plus(+)string(2)' }

      it { should eq 4.0 }
    end

    context do
      let(:string) { '2++3' }

      it { should eq 5.0 }
    end

    context do
      let(:string) { '10----8' }

      it { should eq 2.0 }
    end

    context do
      let(:string) { '3***5' }

      it { should eq 15.0 }
    end

    context do
      let(:string) { '17...42+2...58' }

      it { should eq 20.0 }
    end

    context do
      let(:string) { '-12.75*2.0+++' }

      it { should eq -25.5 }
    end

    context do
      let(:string) { '-12.75*2.0----' }

      it { should eq -25.5 }
    end

    context do
      let(:string) { '-12.75*2.0***' }

      it { should eq -25.5 }
    end

    context do
      let(:string) { '-12.75*2.0...' }

      it { should eq -25.5 }
    end
  end
end
